import java.util.AbstractSet;
import java.util.Iterator;
import java.util.Set;

class IntSparseSet extends AbstractSet<Integer> {
    private int[] dense, sparse;
    private int capacity, low, high, size = 0;

    public IntSparseSet(int low, int high) {
        capacity = high - low;
        this.low = low;
        this.high = high;
        dense = new int[capacity];
        sparse = new int[capacity];
    }

    class IntSetIterator implements Iterator<Integer> {
        int iterSize = 0;

        public boolean hasNext() {
            return iterSize < size;
        }

        public Integer next() {
            return dense[iterSize++] + low;
        }

        public void remove(){
            IntSparseSet.this.remove(dense[iterSize-1]+low);
        }
    }

    @Override
    public Iterator<Integer> iterator() {
        return new IntSetIterator();
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean contains(Object o) {
        int value = (Integer) o - low;
        if (value > capacity || value < 0) {
            return false;
        }

        return value < capacity && sparse[value] < size && dense[sparse[value]] == value;
    }

    private boolean localContains(int a) {
        return a < capacity && sparse[a] < size && dense[sparse[a]] == a;
    }

    @Override
    public boolean remove(Object o) {
        int value = (Integer) o - low;

        if (value > capacity || value < 0) {
            return false;
        }
        if (localContains(value)) {
            dense[sparse[value]] = dense[size - 1];
            sparse[dense[size - 1]] = sparse[value];
            --size;
            return true;
        }
        return false;
    }

    @Override
    public boolean add(Integer value) {
        value -= low;

        if (value > capacity || value < 0) {
            return false;
        }
        if (!localContains(value)) {
            sparse[value] = size;
            dense[size] = value;
            size++;
            return true;
        }
        return false;
    }

    public void clear() {
        size = 0;
    }
}