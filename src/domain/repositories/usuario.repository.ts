import { Usuario } from '../entities/usuario.entity';
import { CreateUsuarioDto, UpdateUsuarioDto } from '../../application/dtos/usuario.dto';

export interface UsuarioRepository {
    findAll(): Promise<Usuario[]>;
    findById(id: string): Promise<Usuario | null>;
    create(data: CreateUsuarioDto): Promise<Usuario>;
    update(id: string, data: UpdateUsuarioDto): Promise<Usuario>;
    delete(id: string): Promise<void>;
}
