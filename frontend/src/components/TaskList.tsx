import React from 'react';
import { Task } from '../api';

interface Props {
  tasks: Task[];
}

const TaskList: React.FC<Props> = ({ tasks }) => {
  return (
    <ul>
      {tasks.map(task => (
        <li key={task.id}>
          <input type="checkbox" checked={task.completed} readOnly />
          {task.title}
        </li>
      ))}
    </ul>
  );
};

export default TaskList;
