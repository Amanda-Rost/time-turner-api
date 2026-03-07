import { Tarefa } from '../entities/tarefa.entity';
import { CreateTarefaDto, UpdateTarefaDto } from '../../application/dtos/tarefa.dto';

export interface TarefaRepository {
    findAll(): Promise<Tarefa[]>;
    findById(id: string): Promise<Tarefa | null>;
    findByUsuarioId(id: string): Promise<Tarefa[] | null>;
    create(data: CreateTarefaDto): Promise<Tarefa>;
    update(id: string, data: UpdateTarefaDto): Promise<Tarefa>;
    delete(id: string): Promise<void>;
}
