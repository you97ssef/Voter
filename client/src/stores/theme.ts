import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'

const THEME_KEY = 'theme'
const HTML_ATTRIBUTE = 'data-theme'

export const DARK_THEME = 'darko'
export const LIGHT_THEME = 'lighto'

export const useThemeStore = defineStore('theme', () => {
    const theme: Ref<string> = ref(loadTheme())

    function toggleTheme(): void {
        const t = theme.value === DARK_THEME ? LIGHT_THEME : DARK_THEME;
        localStorage.setItem(THEME_KEY, t);
        document.documentElement.setAttribute(HTML_ATTRIBUTE, t);
        theme.value = t;
    }

    function loadTheme(): string {
        const t = localStorage.getItem(THEME_KEY) || LIGHT_THEME;
        document.documentElement.setAttribute(HTML_ATTRIBUTE, t)
        return t;
    }

    return { theme, toggleTheme }
})
