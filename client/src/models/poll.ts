import type { Option } from "./option";
import type { Vote } from "./vote";

export interface Poll {
    id: number;
    description: string;
    user_id: number;
    private_code: string | null;
}

export interface PollWithOptions extends Poll {
    options: Option[];
}

export function buildPollsWithOptions(polls: Poll[], options: Option[]): PollWithOptions[] {
    return polls.map(poll => {
        return {
            ...poll,
            options: options.filter(option => option.poll_id === poll.id)
        };
    });
}

export interface Polls {
    polls: Poll[];
    options: Option[];
}

export interface PollResults {
    poll: Poll;
    options: Option[];
    votes: Vote[];
}

export interface Results {
    poll: Poll;
    count: number;
    options: {
        details: Option;
        percentage: number;
    }[];
    votes: Vote[];
}

export function buildResults(poll: PollResults): Results {
    if (poll.votes === null) {
        poll.votes = [];
    }
    
    const count = poll.votes.length;
    const options = poll.options.map(option => {
        return {
            details: option,
            percentage: Math.round((poll.votes.filter(vote => vote.option_id === option.id).length / count) * 100) || 0,
        };
    });
    return {
        poll: poll.poll,
        count,
        options,
        votes: poll.votes
    };
}

export function addVote(results: Results, vote: Vote): Results {
    results.votes.push(vote);
    results.count++;
    results.options = results.options.map(option => {
        return {
            ...option,
            percentage: Math.round((results.votes.filter(vote => vote.option_id === option.details.id).length / results.count) * 100) || 0,
        };
    });
    return results;
}


export interface NewPoll {
    description: string;
    options: string[];
    private: boolean;
}