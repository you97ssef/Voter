export interface Vote {
    id: number;
    user_id: number;
    option_id: number;
    poll_id: number;
    guest: string | null;
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
