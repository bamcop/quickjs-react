import React from 'react'
import QSwitch from './components/qswitch'

export default function App() {
  return (
      <div id="root">
        <h1>Hello, world!</h1>
        <Counter />
        <QSwitch />
      </div>
  );
}

function Counter() {
  const [count, setCount] = React.useState(0);
  return (
      <button onClick={() => setCount(count + 1)}>
        You clicked me {count} times
      </button>
  );
}
