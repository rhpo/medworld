import { API_CONFIG } from './config';
import { transformSnakeToCamel, transformCamelToSnake } from './transforms';

export class ApiError extends Error {
    constructor(
        public status: number,
        public message: string,
        public data?: any
    ) {
        super(message);
        this.name = 'ApiError';
    }
}

export interface ApiResponse<T> {
    data?: T;
    message?: string;
    errors?: Record<string, string[]>;
}

class ApiClient {
    private baseURL: string;

    constructor() {
        this.baseURL = API_CONFIG.BASE_URL;
    }

    private async request<T>(
        endpoint: string,
        options: RequestInit = {}
    ): Promise<T> {
        const url = `${this.baseURL}${endpoint}`;

        // Get auth token from localStorage
        const token = typeof localStorage !== 'undefined' ? localStorage.getItem('authToken') : null;

        const headers: HeadersInit = {
            ...API_CONFIG.HEADERS,
            ...options.headers,
        };

        // Add Authorization header if token exists
        if (token) {
            (headers as any)['Authorization'] = `Bearer ${token}`;
        }

        const config: RequestInit = {
            ...options,
            headers,
            credentials: 'include', // Important for cookie-based auth
        };

        try {
            const response = await fetch(url, config);

            // Handle non-JSON responses
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                if (!response.ok) {
                    throw new ApiError(
                        response.status,
                        `HTTP ${response.status}: ${response.statusText}`
                    );
                }
                return {} as T;
            }

            const data: ApiResponse<T> = await response.json();

            if (!response.ok) {
                throw new ApiError(
                    response.status,
                    data.message || `HTTP ${response.status}: ${response.statusText}`,
                    data.errors || data
                );
            }

            // Laravel returns data in various formats, handle both
            // Direct data or wrapped in {data: ...} or paginated {data: [...], ...}
            let result: any;
            if (data.data !== undefined) {
                result = data.data;
            } else {
                result = data;
            }

            // Transform snake_case keys to camelCase for consistency with frontend types
            return transformSnakeToCamel(result) as T;
        } catch (error) {
            if (error instanceof ApiError) {
                throw error;
            }

            if (error instanceof TypeError) {
                throw new ApiError(0, 'Network error: Unable to connect to server');
            }

            throw new ApiError(500, 'An unexpected error occurred');
        }
    }

    async get<T>(endpoint: string): Promise<T> {
        return this.request<T>(endpoint, {
            method: 'GET',
        });
    }

    async post<T>(endpoint: string, data?: any): Promise<T> {
        const transformedData = data ? transformCamelToSnake(data) : undefined;
        return this.request<T>(endpoint, {
            method: 'POST',
            body: transformedData ? JSON.stringify(transformedData) : undefined,
        });
    }

    async put<T>(endpoint: string, data?: any): Promise<T> {
        const transformedData = data ? transformCamelToSnake(data) : undefined;
        return this.request<T>(endpoint, {
            method: 'PUT',
            body: transformedData ? JSON.stringify(transformedData) : undefined,
        });
    }

    async patch<T>(endpoint: string, data?: any): Promise<T> {
        const transformedData = data ? transformCamelToSnake(data) : undefined;
        return this.request<T>(endpoint, {
            method: 'PATCH',
            body: transformedData ? JSON.stringify(transformedData) : undefined,
        });
    }

    async delete<T>(endpoint: string): Promise<T> {
        return this.request<T>(endpoint, {
            method: 'DELETE',
        });
    }
}

export const apiClient = new ApiClient();
