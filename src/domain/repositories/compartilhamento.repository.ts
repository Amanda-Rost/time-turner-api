import { Compartilhamento } from '../entities/compartilhamento.entity';
import { CreateCompartilhamentoDto, UpdateCompartilhamentoDto } from '../../application/dtos/compartilhamento.dto';

export interface CompartilhamentoRepository {
    findAll(): Promise<Compartilhamento[]>;
    findById(id: string): Promise<Compartilhamento | null>;
    findByTarefaId(id: string): Promise<Compartilhamento[] | null>;
    findByUsuarioId(id: string): Promise<Compartilhamento[] | null>;
    create(data: CreateCompartilhamentoDto): Promise<Compartilhamento>;
    update(id: string, data: UpdateCompartilhamentoDto): Promise<Compartilhamento>;
    delete(id: string): Promise<void>;
}
