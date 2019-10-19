class Bird: 
    def tweet(self):
        print("chuchuchu")

    def fly(self):
        print("papapa")

class Eagle(Bird):
    pass

# 鴕鳥
class Ostrich(Bird):
    def fly(self):
        raise NotImplementedError

if __name__ == '__main__':
    birds = [Bird(),Ostrich() , Eagle()]
    for bird in birds:
        # Error happens 
        # Ostrich can not replace Bird.
        bird.fly() 


