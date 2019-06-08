#include <iostream>
#include <vector>
#include <fstream>
#include <set>
#include <algorithm>

using namespace std;

void get_bigram(const string &word, set<string> &res){
    if(word.size() <=2){
        res.insert(word);
        return;
    }
    for(int i = 0; i < word.size() - 1; i++){
        res.insert(word.substr(i, 2));
    }
}


struct word_stat{
    int freq;
    string word;
    set<string> bigram_set;
    word_stat(const string w, int freq) : freq(freq), word(w){
        get_bigram(word, bigram_set);
    };
};


vector<word_stat> read_freq_di(const string &file_name){
    vector<word_stat> res;
    ifstream in(file_name);
    string word;
    int freq;
    while (in >> word >> freq){ res.emplace_back(word,freq); }
    in.close();
    return res;
}


vector<string> read_input_words(){
    vector<string> res;
    string s;
    while (cin >> s){ res.push_back(s); }
    return res;
}

double ident_mesure(const set<string> &a, const set<string> &b){
    vector<string> u = vector<string>();
    vector<string> i = vector<string>();
    set_intersection(a.begin(),a.end(),b.begin(),b.end(), back_inserter(i));
    if(i.empty()) {
        return 0;
    }
    set_union(a.begin(),a.end(),b.begin(),b.end(), back_inserter(u));
    return (double)i.size() / (double)u.size();
}

string get_right(const vector<word_stat> &freq_di, const vector<string> &words, const string &word){
    word_stat correct_word(word,0);
    double max = 0;
    set<string> word_bigram;
    get_bigram(word, word_bigram);
    for(auto cword : freq_di){
        double k = ident_mesure(word_bigram, cword.bigram_set);
        if(abs(k - 1) < 0.001){
            return cword.word;
        }
        if(k - max > 0.001){
            max = k;
            correct_word = cword;
        } else if (abs(k - max) < 0.001 ){
            if(cword.freq > correct_word.freq){
                max = k;
                correct_word = cword;
            } else if(cword.freq == correct_word.freq && cword.word < correct_word.word ){
                max = k;
                correct_word = cword;
            }
        }
    }
    return correct_word.word;
}

void put_correct(const vector<word_stat> &freq_di, const vector<string> &words, vector<string> &correct){
    for(auto word : words){
        correct.push_back(get_right(freq_di,words, word));
    }
}

int main(){
    vector<word_stat> dict = read_freq_di("count_big.txt");
    vector<string> words = read_input_words();
    vector<string> out;
    put_correct(dict, words, out);
    for(const string& word : out){
        cout << word << endl;
    }
    return 0;
}