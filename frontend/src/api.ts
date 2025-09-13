export interface Task {
  id: number;
  title: string;
  completed: boolean;
}

export async function fetchTasks(): Promise<Task[]> {
  const res = await fetch('http://localhost:8080/tasks');
  if (!res.ok) throw new Error('Failed to fetch tasks');
  return res.json();
}
