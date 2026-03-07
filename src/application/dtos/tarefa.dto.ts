// src/application/dtos/your-entity.dto.ts
import { z } from 'zod';
import { createSchema } from 'zod-openapi';
import { MediaTypeObject } from 'zod-openapi/dist/openapi3-ts/dist/oas30';

// Zod schemas for validation
export const CreateTarefaSchema = z.object({
    // Add your properties here with validation
    name: z.string().min(3).describe('Entity name'),
    // ... other properties
});

export type CreateTarefaDto = z.infer<typeof CreateTarefaSchema>;

export const UpdateTarefaSchema = CreateTarefaSchema.partial();
export type UpdateTarefaDto = z.infer<typeof UpdateTarefaSchema>;

export const ResponseTarefaSchema = CreateTarefaSchema.extend({
    id: z.string(),
    createdAt: z.date(),
    updatedAt: z.date(),
});

export type ResponseTarefaDto = z.infer<typeof ResponseTarefaSchema>;

// OpenAPI schemas for documentation
export const CreateTarefaOpenApi = createSchema(CreateTarefaSchema) as MediaTypeObject;
export const UpdateTarefaOpenApi = createSchema(UpdateTarefaSchema) as MediaTypeObject;
export const ResponseTarefaOpenApi = createSchema(ResponseTarefaSchema) as MediaTypeObject;
