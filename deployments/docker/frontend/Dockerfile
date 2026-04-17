# syntax=docker/dockerfile:1.7

FROM node:22-alpine AS builder

WORKDIR /src/apps/frontend

ARG NEXT_PUBLIC_API_BASE_URL
ENV NEXT_PUBLIC_API_BASE_URL=${NEXT_PUBLIC_API_BASE_URL}

COPY apps/frontend/package.json ./
RUN npm install

COPY apps/frontend ./

RUN npm run build

FROM node:22-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production

COPY --from=builder /src/apps/frontend/.next/standalone ./
COPY --from=builder /src/apps/frontend/.next/static ./.next/static
COPY --from=builder /src/apps/frontend/public ./public

EXPOSE 3000

CMD ["node", "server.js"]
