import React, { useEffect, useState } from 'react';
import TaskList from './components/TaskList';
import { fetchTasks } from './api';

const App: React.FC = () => {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetchTasks().then(setTasks);
  }, []);

  return (
    <div style={{ padding: '2rem' }}>
      <h1>📝 Task Tracker</h1>
      <TaskList tasks={tasks} />
    </div>
  );
};

export default App;
