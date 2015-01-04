#import <map>

class Plugboard {

public:
    Plugboard(const char *plugboard_file);
    char translate(char c);

private:
    std::map<char,char> translation_map;

};