import _deepmerge from "deepmerge";

export function deepmerge<T>(x: Partial<T>, y: Partial<T>): T;
export function deepmerge<T1, T2>(x: Partial<T1>, y: Partial<T2>): T1 & T2;
export function deepmerge<T1, T2>(x: Partial<T1>, y: Partial<T2>): unknown {
  return _deepmerge(x, y, {
    isMergeableObject: isPlainObject,
  });
}

export function isPlainObject(o: unknown): boolean {
  if (!isObject(o)) return false;

  const ctor = o.constructor;
  if (ctor === undefined) return true;

  const prot = ctor.prototype;
  if (!isObject(prot)) return false;

  return Object.prototype.hasOwnProperty.call(prot, "isPrototypeOf");
}

function isObject(o: unknown): o is object {
  return Object.prototype.toString.call(o) === "[object Object]";
}
