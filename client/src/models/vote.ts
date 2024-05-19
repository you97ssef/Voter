export interface Vote {
    id: number;
    user_id: number;
    option_id: number;
    poll_id: number;
    guest: string | null;
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
        }

        return vote;
    });
}
