export enum PlanID {
    premium = 'premium',
    standard = 'standard',
    basic = 'basic'
}

export type Plan = {
    id: number;
    planID: PlanID;
    name: string;
    minPlan: number;
}
