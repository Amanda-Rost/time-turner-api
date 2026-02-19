export enum Status {
  planejada = 'planejada',
  concluida = '',
  adiada = 'adiada',
  cancelada = 'cancelada',
}

export enum Prioridade {
  baixa = 'baixa',
  media = 'media',
  alta = 'alta',
}

export class Tarefa {
    id: string;
    usuarioId: string;
    titulo:string;
    data: Date;
    alarme: boolean;
    prioridade: Prioridade;
    status:Status;
    descricao?: string;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;

    constructor(props: Omit<Tarefa, 'id' | 'createdAt' | 'updatedAt'>, id?: string) {
        Object.assign(this, props);
        this.id = id ?? crypto.randomUUID();
        this.createdAt = new Date();
        this.updatedAt = new Date();
    }
}