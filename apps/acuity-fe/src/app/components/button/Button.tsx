import classNames from 'classnames';
import { ButtonHTMLAttributes, ReactNode } from 'react';
import { SyncLoader } from 'react-spinners';
import styles from './Button.module.scss';

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  children: ReactNode;
  isLoading?: boolean;
}

export function Button({ children, isLoading = false, ...props }: ButtonProps) {
  return (
    <button
      {...props}
      className={classNames(styles.root, isLoading && styles.loading)}
      disabled={isLoading || props.disabled}
    >
      {isLoading && <SyncLoader color="#67599e" loading={isLoading} size={3} />}
      {children}
    </button>
  );
}
