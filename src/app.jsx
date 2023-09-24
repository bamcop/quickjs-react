import React from 'react'
import QSwitch from './components/qswitch'
import QMenu from './components/qmenu'
import QSelect from './components/qselect'
import QAutocomplete from './components/qautocomplete'
import QDisclosure from './components/qdisclosure'
import QPopover from './components/qpopover'
import QRadioGroup from './components/qradio_group'
import QTabs from './components/qtabs'
import QTransition from './components/qtransition'

export default function App() {
  return (
    <div id="root">
      <div className='grid grid-cols-3 bg-white gap-4 p-4'>
        <div className='bg-gray-400 p-4'><QMenu /></div>
        <div className='bg-gray-400 p-4'><QSelect /></div>
        <div className='bg-gray-400 p-4'><QAutocomplete /></div>
        <div className='bg-gray-400 p-4'><QSwitch /></div>
        <div className='bg-gray-400 p-4'><QPopover /></div>
        <div className='bg-gray-400 p-4'><QDisclosure /></div>
        <div className='bg-gray-400 p-4'><QRadioGroup /></div>
        <div className='bg-gray-400 p-4'><QTabs /></div>
        <div className='bg-gray-400 p-4'><QTransition /></div>
        <div className='bg-gray-400 p-4'><Counter /></div>
      </div>
    </div>
  );
}

function Counter() {
  const [count, setCount] = React.useState(0);

  return (
    <button className='bg-gray-300' onClick={() => setCount(count + 1)}>
      You clicked me <span className='text-red-500'>{count}</span> times
    </button>
  );
}
