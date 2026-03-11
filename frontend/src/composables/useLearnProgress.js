import { ref } from 'vue'

const STORAGE_KEY = '110claw_completed_days'

function load() {
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
  } catch {
    return []
  }
}

const completedDays = ref(load())

export function useLearnProgress() {
  function isCompleted(day) {
    return completedDays.value.includes(Number(day))
  }

  function toggleDay(day) {
    const d = Number(day)
    if (completedDays.value.includes(d)) {
      completedDays.value = completedDays.value.filter(x => x !== d)
    } else {
      completedDays.value = [...completedDays.value, d]
    }
    localStorage.setItem(STORAGE_KEY, JSON.stringify(completedDays.value))
  }

  function reset() {
    completedDays.value = []
    localStorage.removeItem(STORAGE_KEY)
  }

  return { completedDays, isCompleted, toggleDay, reset }
}
