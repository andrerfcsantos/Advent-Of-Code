interface Node<T> {
  key: number;
  value: T;
}

export class PriorityQueue<T> {
  private heap: Node<T>[] = [];

  parent(index: number): number {
    return Math.floor((index - 1) / 2);
  }

  left(index: number): number {
    return 2 * index + 1;
  }

  right(index: number): number {
    return 2 * index + 2;
  }

  hasLeft(index: number): boolean {
    return this.left(index) < this.heap.length;
  }

  hasRight(index: number): boolean {
    return this.right(index) < this.heap.length;
  }

  isEmpty(): boolean {
    return this.heap.length == 0;
  }

  peek(): T | null {
    return this.heap.length == 0 ? null : this.heap[0].value;
  }

  size(): number {
    return this.heap.length;
  }

  has(element: T, equalsFunc?: (a: T, b: T) => boolean): boolean {
    if (equalsFunc) {
      return this.heap.some((n) => equalsFunc(n.value, element));
    }

    return this.heap.some((n) => n.value == element);
  }

  pop(): T | null {
    if (this.heap.length == 0) return null;

    this.swap(0, this.heap.length - 1);
    const item = this.heap.pop();
    this.shiftDown(0);

    return item?.value ?? null;
  }

  private shiftDown(from: number) {
    let current = from;
    while (this.hasLeft(current)) {
      let smallerChild = this.left(current);
      if (
        this.hasRight(current) &&
        this.heap[this.right(current)].key < this.heap[this.left(current)].key
      )
        smallerChild = this.right(current);

      if (this.heap[smallerChild].key > this.heap[current].key) break;

      this.swap(current, smallerChild);
      current = smallerChild;
    }
  }

  private shiftUp(from: number, maxTo = 0) {
    const item = this.heap[from];
    let i = from;
    while (i > maxTo) {
      const p = this.parent(i);
      if (this.heap[p].key < this.heap[i].key) break;
      this.swap(i, p);
      i = p;
    }
    this.heap[i] = item;
  }

  insert(item: T, priority: number) {
    this.heap.push({ key: priority, value: item });
    this.shiftUp(this.heap.length - 1);
  }

  replace(item: T, newPriority: number, equalsFunc: (a: T, b: T) => boolean) {
    const index = this.heap.findIndex((n) => equalsFunc(n.value, item));
    if (index == -1) return;

    const oldPriority = this.heap[index].key;
    this.heap[index].key = newPriority;
    if (oldPriority < newPriority) {
      this.shiftDown(index);
    } else {
      this.shiftUp(index);
    }
  }

  private swap(a: number, b: number) {
    const tmp = this.heap[a];
    this.heap[a] = this.heap[b];
    this.heap[b] = tmp;
  }
}
