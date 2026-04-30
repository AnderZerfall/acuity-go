import {
  Control,
  Controller,
  FieldError,
  FieldValues,
  Path,
} from 'react-hook-form';
import styles from './URLField.module.scss';

interface URLFieldProps<T extends Omit<FieldValues, 'error'>> {
  name: Path<T>;
  control: Control<T>;
  label?: string;
  placeholder?: string;
  rules?: object;
  error?: FieldError;
}

export function URLField<T extends FieldValues>({
  name,
  control,
  label,
  placeholder,
  rules,
}: URLFieldProps<T>) {
  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState: { error } }) => (
        <div
          className={`${styles.inputContainer} ${error ? styles.hasError : ''}`}
        >
          {label && <span className={styles.label}>{label}</span>}

          <div className={styles.wrapper}>
            <input
              {...field}
              type="text"
              placeholder={placeholder}
              autoComplete="off"
            />
          </div>

          <div className={styles.errorMessage}>
            {error && <span>{error.message}</span>}
          </div>
        </div>
      )}
    />
  );
}
