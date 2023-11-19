export interface Artist {
    ID: number
    CreatedAt?: string
    name: string
    contact: string
    genere: string
    description?: string
    socials: Social[]
}

export interface Social {
    platform: string
    link: string
}