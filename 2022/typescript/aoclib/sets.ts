export function intersectSets<T>(...sets: Set<T>[]): Set<T> {
  const result = new Set<T>();
  for (const set of sets) {
    for (const item of set) {
      if (sets.every((s) => s.has(item))) {
        result.add(item);
      }
    }
  }
  return result;
}
