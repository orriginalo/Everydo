export function getTasksByTypes(tasks) {
  const result = {
    daily: [],
    weekly: [],
    custom: [],
  }

  console.log('GetTasksByTypes')
  console.log(tasks)

  for (const task of tasks) {
    const type = task.reload_type

    if (result[type]) {
      result[type].push(task)
    }
  }

  return result
}
