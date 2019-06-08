#include <iostream>
#include "supercalc.hpp"

using namespace std;

int main()
{
	SuperCalc<int> sc(5, 5);
	for (int j = 1; j < 5; j++) sc(0, j) = sc(0, j-1) + 1;
	for (int i = 1; i < 5; i++) {
		for (int j = 0; j < 5; j++) sc(i, j) = sc(i-1, j) + 1;
	}
	
	sc(0, 0) = 0;
	for (int i = 0; i < 5; i++) {
		for (int j = 0; j < 5; j++) cout << (int)sc(i, j) << " ";
		cout << endl;
	}
	return 0;
}