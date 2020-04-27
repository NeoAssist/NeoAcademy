FROM node:14-alpine as base

FROM base as builder

WORKDIR /root

COPY . .

RUN yarn install

FROM base as runner

WORKDIR /root

COPY --from=builder /root /root

EXPOSE 80

CMD ["yarn", "dev"]
