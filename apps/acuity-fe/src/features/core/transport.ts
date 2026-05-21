import { createConnectTransport } from '@connectrpc/connect-web';

const BASE_URL = import.meta.env.VITE_BASE_URL;

export const transport = createConnectTransport({
  baseUrl: BASE_URL,
});
