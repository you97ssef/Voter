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
                    primary: '#F5F5F5',
                    'primary-content': '#0A0A0A',
                    secondary: '#A3A3A3',
                    'secondary-content': '#0A0A0A',
                    accent: '#474747',
                    'accent-content': '#F5F5F5',

                    neutral: '#141414',

                    'base-100': '#181818',
                    'base-200': '#101010',
                    'base-300': '#000000',
                    
                    info: '#3876BF',
                    'info-content': '#050A10',
                    success: '#00f08c',
                    'success-content': '#00140C',
                    warning: '#ffa600',
                    'warning-content': '#140D00',
                    error: '#ff0037',
                    'error-content': '#140004',

                    '--rounded-box': '1.5rem',
                    '--rounded-btn': '1rem',
                    '--rounded-badge': '2rem',

                    '--animation-btn': '.25s',
                    '--animation-input': '.2s',

                    '--btn-text-case': 'uppercase',
                    '--navbar-padding': '1rem',
                    '--border-btn': '2px'
                }, 
                lighto: {
                    primary: '#0A0A0A',
                    'primary-content': '#F5F5F5',
                    secondary: '#1F1F1F',
                    'secondary-content': '#F5F5F5',
                    accent: '#707070',
                    'accent-content': '#0A0A0A',

                    neutral: '#F5F5F5',

                    'base-100': '#D6D6D6',
                    'base-200': '#EBEBEB',
                    'base-300': '#FFFFFF',
                    
                    info: '#29588e',
                    'info-content': '#F5F5F5',
                    success: '#008F53',
                    'success-content': '#F5F5F5',
                    warning: '#8f5d00',
                    'warning-content': '#F5F5F5',
                    error: '#A30023',
                    'error-content': '#F5F5F5',

                    '--rounded-box': '1.5rem',
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
