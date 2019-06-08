public class Element<T> {
    private Element<T> parent;
    private T value;
    private int rank;

    public Element(T x) {
        value = x;
        rank = 0;
        parent = this;
    }

    public T x() {
        return value;
    }

    public boolean equivalent(Element<T> elem) {
        return this.find() == elem.find();
    }

    public void union(Element<T> elem) {
        Element<T> x = this.find();
        Element<T> y = elem.find();
        if(x.rank < y.rank){
            x.parent = y;
        }else{
            y.parent = x;
            if (x.rank == y.rank){
                x.rank++;
            }
        }
    }

    public Element<T> find(){
        if (this == this.parent){
            return this;
        }else{
            return this.parent = this.parent.find();
        }
    }
}