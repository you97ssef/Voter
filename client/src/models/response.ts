export interface OkResponse<T> {
    message: string
    data: T
}

export interface CreatedResponse<T> {
    message: string
    data: T
}

export interface ErrorResponse {
    message: string
}
