#include <iostream>
#include <vector>
#include <functional>

using namespace std;


template <class T>
class Cell{

public:
    function<T()> f_value;
    Cell():f_value(){};
    Cell(T val) : f_value([val]()-> T { return val;}) {};
    Cell(function<T()> (formula)) : f_value(formula){};

    Cell<T>&operator=(T val){
        f_value = [val]()-> T { return val;};
        return *this;
    }
    Cell<T>&operator=(function<T()> f_val){
        f_value = f_val;
        return *this;
    }

    Cell<T> operator-() const{
        return Cell([this]() -> T { return -f_value();});
    }

    Cell<T> operator+() const{
        return Cell([this]() -> T { return f_value();});
    }

    Cell<T> operator+(const Cell<T> &obj) const{
        return Cell([&obj, this]() -> T { return f_value() + obj.f_value();});
    }

    Cell<T> operator-(const Cell<T> &obj) const{
        return Cell([&obj, this]() -> T { return f_value() - obj.f_value();});
    }

    Cell<T> operator*(const Cell<T> &obj) const{
        return Cell([&obj, this]() -> T { return f_value() * obj.f_value();});
    }

    Cell<T> operator/(const Cell<T> &obj) const{
        return Cell([&obj, this]() -> T { return f_value() / obj.f_value();});
    }

    Cell<T> operator+=(const Cell<T> &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() + obj.f_value();};
    }

    Cell<T> operator-=(const Cell<T> &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() - obj.f_value();};
    }

    Cell<T> operator*=(const Cell<T> &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() * obj.f_value();};
    }

    Cell<T> operator/=(const Cell<T> &obj){
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() / obj.f_value();};
    }




    //-------------------------------------------------------------------


    Cell<T> operator+(const T &obj) const{
        return Cell([&obj, this]() -> T { return f_value() + obj;});
    }

    Cell<T> operator-(const T &obj) const{
        return Cell([&obj, this]() -> T { return f_value() - obj;});
    }

    Cell<T> operator*(const T &obj) const{
        return Cell([&obj, this]() -> T { return f_value() * obj;});
    }

    Cell<T> operator/(const T &obj) const{
        return Cell([&obj, this]() -> T { return f_value() / obj;});
    }

    Cell<T> operator+=(const T &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() + obj;};
    }

    Cell<T> operator-=(const T &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() - obj;};
    }

    Cell<T> operator*=(const T &obj) {
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() * obj;};
    }

    Cell<T> operator/=(const T &obj){
        Cell<T> tmp = *this;
        return f_value = [&obj, tmp]() -> T { return tmp.f_value() / obj;};
    }





    explicit operator T(){
        return f_value();
    }
};

template <class T>
class SuperCalc {
private:
    vector<vector<Cell<T>>> data;
public:
    SuperCalc(int m, int n){
        data = vector<vector<Cell<T>>>(m);
        for(int i = 0; i < m; i++){
            data[i] = vector<Cell<T>>(n);
        }
    };
    Cell<T>& operator() (int i, int j){
        return data.at(i).at(j);
    };
};

template <class T>
Cell<T> operator-(int i, Cell<T> &obj) {
    return Cell<T>([i, &obj]() -> T { return i - obj.f_value();});
}

template <class T>
Cell<T> operator+(int i, Cell<T> &obj) {
    return Cell<T>([i, &obj]() -> T { return i + obj.f_value();});
}

template <class T>
Cell<T> operator*(int i, Cell<T> &obj) {
    return Cell<T>([i, &obj]() -> T { return i * obj.f_value();});
}

template <class T>
Cell<T> operator/(int i, Cell<T> &obj) {
    return Cell<T>([i, &obj]() -> T { return i / obj.f_value();});
}