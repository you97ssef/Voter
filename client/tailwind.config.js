/** @type {import('tailwindcss').Config} */
export default {
    content: ['./index.html', './src/**/*.{vue,ts}'],
    theme: {
        extend: {
            fontFamily: {
                sans: ['abeezee', 'sans-serif'],
                lobster: ['lobster-two', 'sans-serif'],
                courgette: ['courgette', 'sans-serif']
            }
        }
    },
    daisyui: {
        themes: [
            {
                darko: {
                    primary: '#000000',
                    'primary-content': '#F5F5F5',
                    secondary: '#0A0A0A',
                    'secondary-content': '#F5F5F5',
                    accent: '#474747',
                    'accent-content': '#F5F5F5',

                    neutral: '#141414',

                    'base-100': '#333333',
                    'base-200': '#252525',
                    'base-300': '#1F1F1F',
                    
                    info: '#3876BF',
                    'info-content': '#050A10',
                    success: '#00f08c',
                    'success-content': '#00140C',
                    warning: '#ffa600',
                    'warning-content': '#140D00',
                    error: '#ff0037',
                    'error-content': '#140004',

                    '--rounded-box': '2rem',
                    '--rounded-btn': '1rem',
                    '--rounded-badge': '2rem',

                    '--animation-btn': '.25s',
                    '--animation-input': '.2s',

                    '--btn-text-case': 'uppercase',
                    '--navbar-padding': '1rem',
                    '--border-btn': '2px'
                }, 
                lighto: {
                    primary: '#FFFFFF',
                    'primary-content': '#0A0A0A',
                    secondary: '#EBEBEB',
                    'secondary-content': '#0A0A0A',
                    accent: '#D6D6D6',
                    'accent-content': '#0A0A0A',

                    neutral: '#F5F5F5',

                    'base-100': '#C2C2C2',
                    'base-200': '#ADADAD',
                    'base-300': '#999999',
                    
                    info: '#29588e',
                    'info-content': '#F5F5F5',
                    success: '#008F53',
                    'success-content': '#F5F5F5',
                    warning: '#8f5d00',
                    'warning-content': '#F5F5F5',
                    error: '#A30023',
                    'error-content': '#F5F5F5',

                    '--rounded-box': '2rem',
                    '--rounded-btn': '1rem',
                    '--rounded-badge': '2rem',

                    '--animation-btn': '.25s',
                    '--animation-input': '.2s',

                    '--btn-text-case': 'uppercase',
                    '--navbar-padding': '1rem',
                    '--border-btn': '2px'
                }
            }
        ]
    },
    plugins: [require('daisyui')]
}
