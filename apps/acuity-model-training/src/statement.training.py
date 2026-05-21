import argparse
import os
from transformers import (
    AutoTokenizer,
    AutoModelForSeq2SeqLM,
    DataCollatorForSeq2Seq,
    Seq2SeqTrainingArguments,
    Seq2SeqTrainer
)
from datasets import Dataset

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--output_dir", type=str, default="../../apps/acuity-be/external/statement")
    args = parser.parse_args()

    model_name = "google-t5/t5-small"

    tokenizer = AutoTokenizer.from_pretrained(model_name)
    model = AutoModelForSeq2SeqLM.from_pretrained(model_name)

    data = {
        "input_text": [
            "summarize_claim: Yesterday on social media people started sharing information that our planet is flat and scientists hide it.",
            "summarize_claim: New medical reports claim that drinking 5 liters of coffee daily grants complete immunity to all known viruses."
        ],
        "target_text": [
            "Earth is flat",
            "Coffee grants virus immunity"
        ]
    }
    dataset = Dataset.from_dict(data)

    def preprocess_function(examples):
        model_inputs = tokenizer(examples["input_text"], max_length=128, truncation=True)
        labels = tokenizer(text_target=examples["target_text"], max_length=16, truncation=True)
        model_inputs["labels"] = labels["input_ids"]
        return model_inputs

    tokenized_dataset = dataset.map(preprocess_function, batched=True)
    data_collator = DataCollatorForSeq2Seq(tokenizer, model=model)

    training_args = Seq2SeqTrainingArguments(
        output_dir="./results",
        eval_strategy="no",
        learning_rate=5e-5,
        per_device_train_batch_size=2,
        weight_decay=0.01,
        save_total_limit=1,
        num_train_epochs=3,
        predict_with_generate=True,
        fp16=True,
        report_to="none"
    )

    trainer = Seq2SeqTrainer(
        model=model,
        args=training_args,
        train_dataset=tokenized_dataset,
        tokenizer=tokenizer,
        data_collator=data_collator,
    )

    print("=== Fine Tunning T5... ===")
    trainer.train()

    print(f"=== Save model to {args.output_dir} ===")
    os.makedirs(args.output_dir, exist_ok=True)

    model.save_pretrained(args.output_dir, safe_serialization=False)
    tokenizer.save_pretrained(args.output_dir)
    print("=== Success ===")

if __name__ == "__main__":
    main()
