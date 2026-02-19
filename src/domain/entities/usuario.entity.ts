export class Usuario {
    id: string;
    nome: string;
    sobrenome?: string;
    email: string;
    senha: string;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;

    constructor(props: Omit<Usuario, 'id' | 'createdAt' | 'updatedAt'>, id?: string) {
        Object.assign(this, props);
        this.id = id ?? crypto.randomUUID();
        this.createdAt = new Date();
        this.updatedAt = new Date();
    }
}