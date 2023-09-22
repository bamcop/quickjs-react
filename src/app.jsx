import React from 'react'

export default function App() {
  return (
      <div id="root" className='bg-green-500'>
        <h1>Hello, world!</h1>
        <Counter />
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
