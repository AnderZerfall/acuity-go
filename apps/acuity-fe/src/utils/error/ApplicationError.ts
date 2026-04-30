export interface ApplicationError extends Error {
  details?: Record<string, string[]>;
}

export const ApplicationError = {
  create: (init: Partial<ApplicationError> = {}): ApplicationError => ({
    ...init,
    name: init.name ?? '',
    details: init.details ?? {},
    message: init.message ?? '',
  }),
};
