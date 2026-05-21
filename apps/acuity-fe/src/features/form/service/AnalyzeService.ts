import { createClient } from "@connectrpc/connect";
import { transport } from "../../core/transport";

export function createAnalyzeService() {
  return createClient(ElizaService, transport);
}
