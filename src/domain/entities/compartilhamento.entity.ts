export enum Permissao {
  leitor = 'leitor',
  participante = 'participante',
  responsavel = 'responsavel',
}

export class Compartilhamento {
    id: string;
    tarefaId: string;
    usuarioId: string;
    permissao: Permissao;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;

    constructor(props: Omit<Compartilhamento, 'id' | 'createdAt' | 'updatedAt'>, id?: string) {
        Object.assign(this, props);
        this.id = id ?? crypto.randomUUID();
        this.createdAt = new Date();
        this.updatedAt = new Date();
    }
}