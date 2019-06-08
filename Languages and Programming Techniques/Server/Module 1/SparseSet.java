import java.util.*;     

public class SparseSet<T extends Hintable> extends AbstractSet<T> {
    ArrayList<T> dense;
    private int size = 0;

    public SparseSet() {
        dense = new ArrayList<>();
    }

    class IntSetIterator implements Iterator<T> {
        int iterSize = 0;

        public boolean hasNext() {
            return iterSize < size;
        }

        public T next() {
            return dense.get(iterSize++);
        }

        public void remove() {
            SparseSet.this.remove(dense.get(iterSize - 1));
        }
    }

    @Override
    public Iterator<T> iterator() {
        return new IntSetIterator();
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean contains(Object obj) {
        T value = (T) obj;
        return dense.get(value.hint()) == value && value.hint() < size;
    }

    public boolean contains(T value) {
        return dense.get(value.hint()) == value && value.hint() < size;
    }

    @Override
    public boolean remove(Object obj) {
        T value = (T) obj;
        if (contains(value)) {
            dense.get(--size).setHint(value.hint());
            dense.set(value.hint(), dense.get(size));
            return true;
        }
        return false;
    }

    @Override
    public boolean add(T value) {
        //order of terms is important
        if (value.hint() >= size || dense.get(value.hint()) != value) {
            dense.add(value);
            value.setHint(size++);
            return true;
        }
        return false;
    }

    public void clear() {
        size = 0;
    }
}