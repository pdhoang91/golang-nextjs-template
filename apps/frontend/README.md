# Web App

Frontend dùng Next.js App Router + TypeScript.

## Cấu trúc chính

- `app/`: route entry và layout
- `features/`: logic và UI theo từng feature
- `services/`: API layer
- `hooks/`: reusable hooks
- `components/`: shared UI/layout
- `types/`: shared type
- `config/`: env config
- `constants/`: route constants

## Chạy local

```bash
cp .env.example .env.local
npm install
npm run dev
```

## Mục tiêu của cấu trúc này

- page không bị phình to
- logic fetch tách khỏi component hiển thị
- shared component và shared service dùng lại được
- dễ scale thêm feature mới

## Luồng dữ liệu

```text
page -> feature component -> feature hook -> service -> api client -> backend
```

## Feature mẫu

- health status
- todo list
- create todo
