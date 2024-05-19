export interface Vote {
    id: number;
    user_id: number;
    option_id: number;
    guest: string | null;
    poll_id: number;
    
    timestamp: number;
    hash: string;
    prev_hash: string;

    valid: boolean | undefined;
}

export interface NewVote {
    option_id: number;
    poll_id: number;
    poll_code: string | null;
}

export interface NewGuestVote {
    option_id: number;
    poll_id: number;
    poll_code: string | null;
    guest: string;
}

export interface ValidatedVote {
    valid: boolean;
    id: number;
}

export function buildValidations(votes: Vote[], validations: ValidatedVote[] ): Vote[] {
    return votes.map(vote => {
        const validation = validations.find(validation => validation.id === vote.id);

        if (validation) {
            vote.valid = validation.valid;
        } else {
            vote.valid = false;
        }

        return vote;
    });
}

export function isVote(vote: any): boolean {
    return vote.id !== undefined
        && vote.user_id !== undefined
        && vote.option_id !== undefined
        && vote.poll_id !== undefined
        && vote.guest !== undefined
        && vote.valid !== undefined
        && vote.timestamp !== undefined
        && vote.hash !== undefined
        && vote.prev_hash !== undefined;
}