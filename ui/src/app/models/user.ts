export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  username: string
  type: number
  name: string
  address: string
  zip: number
  city: string
}

export interface AuthToken {
  exp: number
  iat: number
  sub: string
  name: string
  type: number
}