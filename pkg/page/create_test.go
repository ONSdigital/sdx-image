package page

import (
	"sdxImage/pkg/model"
	"sdxImage/pkg/test"
	"testing"
)

func TestCreateMbs(t *testing.T) {
	test.SetCwdToRoot()

	survey := &model.Survey{
		Title:       "Monthly Business Survey",
		SurveyId:    "009",
		FormType:    "0106",
		Respondent:  "12346789012A",
		SubmittedAt: "13 January 2023 16:38:46",
		Sections: []*model.Section{
			{
				Title: "",
				Instances: []*model.Instance{
					{
						Id: 0,
						Answers: []*model.Answer{
							{
								Text:  "Are you able to report for the period from 01/05/2016 to 01/06/2016?",
								QCode: "9999",
								Value: "Yes, I can report for this period",
							},
							{
								Text:  "Total turnover excluding VAT?",
								QCode: "40",
								Value: "56000",
							},
							{
								Text:  "Please explain any changes in your turnover figures from the previous return, if applicable.",
								QCode: "146",
								Value: "No changes!",
							},
						},
					},
				},
			},
		},
	}

	result := Create(survey)
	err := test.SaveJPG("temp/mbs-test.jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for mbs_0106 with error: %q", err.Error())
	}
}

func TestCreateInstances(t *testing.T) {
	test.SetCwdToRoot()

	survey := &model.Survey{
		Title:       "Instances Test",
		SurveyId:    "123",
		FormType:    "0106",
		Respondent:  "12346789012A",
		SubmittedAt: "13 January 2023 16:38:46",
		Sections: []*model.Section{
			{
				Title: "No Instances",
				Instances: []*model.Instance{
					{
						Id: 0,
						Answers: []*model.Answer{
							{
								Text:  "What was the business's total expenditure?",
								QCode: "200",
								Value: "10000",
							},
						},
					},
				},
			},
			{
				Title: "Three Instances",
				Instances: []*model.Instance{
					{
						Id: 1,
						Answers: []*model.Answer{
							{
								Text:  "What was the business's expenditure on employment costs for the period?",
								QCode: "450",
								Value: "2000",
							},
							{
								Text:  "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								QCode: "451",
								Value: "1000",
							},
						},
					},
					{
						Id: 2,
						Answers: []*model.Answer{
							{
								Text:  "What was the business's expenditure on employment costs for the period?",
								QCode: "450",
								Value: "3000",
							},
							{
								Text:  "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								QCode: "451",
								Value: "1000",
							},
						},
					},
					{
						Id: 3,
						Answers: []*model.Answer{
							{
								Text:  "What was the business's expenditure on employment costs for the period?",
								QCode: "450",
								Value: "2000",
							},
							{
								Text:  "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								QCode: "451",
								Value: "1000",
							},
						},
					},
				},
			},
		},
	}

	result := Create(survey)
	err := test.SaveJPG("temp/instance-test.jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for instance test with error: %q", err.Error())
	}
}
