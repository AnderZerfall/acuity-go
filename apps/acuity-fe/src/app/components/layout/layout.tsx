import { ReactNode } from 'react';
import styles from './Layout.module.scss';

export function Layout({ children }: { children: ReactNode }) {
  return (
    <main className={styles.root}>
      <div className={styles.container}>{children}</div>
    </main>
  );
}
