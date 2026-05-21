import argparse
import os
from transformers import (
    AutoTokenizer,
    AutoModelForSequenceClassification,
    TrainingArguments,
    Trainer,
    DataCollatorWithPadding
)
from datasets import Dataset

def main():
  model_dir = "../../apps/acuity-be/external/nlp-model"

  if not os.path.exists(model_dir):
    raise FileNotFoundError(f"Model is not found: {model_dir}")

    parser = argparse.ArgumentParser()
    parser.add_argument("--output_dir", type=str, default="../../apps/acuity-be/external/emotional-analysis")
    args = parser.parse_args()

    tokenizer = AutoTokenizer.from_pretrained(model_dir)
    model = AutoModelForSeq2SeqLM.from_pretrained(model_dir)

    data = {
        "text": [
            "I hate everything about this project, it makes me so mad!",
            "Wow, I absolutely love the new features of this app!",
            "Oh brilliant, another error to debug on a Friday night."
        ],
        "label": [1, 5, 12]
    }
    dataset = Dataset.from_dict(data)

    def preprocess_function(examples):
        return tokenizer(examples["text"], truncation=True, max_length=128)

    tokenized_dataset = dataset.map(preprocess_function, batched=True)
    data_collator = DataCollatorWithPadding(tokenizer, model=model)

    training_args = TrainingArguments(
        output_dir="./results_emotion",
        eval_strategy="no",
        learning_rate=2e-5,
        per_device_train_batch_size=4,
        num_train_epochs=3,
        weight_decay=0.01,
        save_total_limit=1,
        fp16=torch.cuda.is_available(),
        report_to="none"
    )

    trainer = Trainer(
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
