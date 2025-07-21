class Student:
    def __init__(self, student_id, academic_score, activity_score):
        self.id = student_id
        self.academic_score = academic_score
        self.activity_score = activity_score


def is_excellent(student):
    if student.academic_score + student.activity_score <= 140:
        return False
    if student.academic_score * 0.7 + student.activity_score * 0.3 < 80:
        return False
    return True


def main():
    n = int(input())

    for _ in range(n):
        student_id, academic, activity = map(int, input().split())

        student = Student(student_id, academic, activity)
        if is_excellent(student):
            print("Excellent")
        else:
            print("Not excellent")


if __name__ == "__main__":
    main()
