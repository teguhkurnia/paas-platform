FROM node:24-alpine

# enable pnpm
RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

# Salin package.json untuk menginstal dependensi
COPY package.json package-lock.json* ./
RUN pnpm install

# Salin sisa kode
COPY . .

# Jalankan server development
EXPOSE 3000
CMD ["pnpm", "dev"]
