import React from 'react'
import { renderToString } from 'react-dom/server'
import App from './app'

const RenderApp = () => {
  return renderToString(<App />)
}

globalThis.React = React
globalThis.RenderApp = RenderApp
