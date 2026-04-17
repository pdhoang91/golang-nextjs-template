export type Todo = {
  id: string;
  title: string;
  description: string;
  completed: boolean;
  created_at: string;
  updated_at: string;
};

export type CreateTodoPayload = {
  title: string;
  description: string;
};
