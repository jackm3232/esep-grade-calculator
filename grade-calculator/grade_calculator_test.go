package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeA2(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 78, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 39, Assignment)
	gradeCalculator.AddGrade("exam 1", 42, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 25, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment', got '%s' instead", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam', got '%s' instead", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay', got '%s' instead", Essay.String())
	}
}

func TestGetPassingGrade(t *testing.T) {
	expected_result := "Pass"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.passFailMode = true

	gradeCalculator.AddGrade("open source assignment", 89, Assignment)
	gradeCalculator.AddGrade("exam 1", 73, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_result := gradeCalculator.GetFinalGrade()

	if expected_result != actual_result {
		t.Errorf("Expected 'Pass'; got '%s' instead", actual_result)
	}
}

func TestGetFailingGrade(t *testing.T) {
	expected_result := "Fail"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.passFailMode = true

	gradeCalculator.AddGrade("open source assignment", 46, Assignment)
	gradeCalculator.AddGrade("exam 1", 26, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 49, Essay)

	actual_result := gradeCalculator.GetFinalGrade()

	if expected_result != actual_result {
		t.Errorf("Expected 'Fail'; got '%s' instead", actual_result)
	}
}
