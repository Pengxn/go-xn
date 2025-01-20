import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import './tailwind.css'

import App from './src/app.jsx'

createRoot(document.getElementById('root')).render(
    <StrictMode>
        <App />
    </StrictMode>,
)

/**
 * Print information on the console.
 * 'APP_NAME', 'APP_VERSION' and 'COMMIT_HASH' is defined in webpack config file.
 */
console.info(`%c${ APP_NAME }%c
Version: ${ APP_VERSION }-${ COMMIT_HASH }
Copyright © ${new Date().getFullYear()}`,
    'font-size: 18px; color: #3b3e43;',
    'font-size: 12px; color: rgba(0,0,0,0.38);');
