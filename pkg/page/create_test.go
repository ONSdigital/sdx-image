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
						Questions: []*model.Question{
							{
								Title: "Are you able to report for the period from 01/05/2016 to 01/06/2016?",
								Answers: []*model.Answer{
									{Type: "Radio", QCode: "9999", Label: "label", Value: "Yes, I can report for this period"},
								},
							},
							{
								Title: "What are the dates of the period that you will be reporting for?",
								Answers: []*model.Answer{
									{Type: "Date", QCode: "11", Label: "From"},
									{Type: "Date", QCode: "12", Label: "To"},
								},
							},
							{
								Title: "For the period, what was the business's total turnover, excluding VAT?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "40", Label: "Total turnover excluding VAT", Value: "56000"},
								},
							},
							{
								Title: "Please explain any changes in your turnover figures from the previous return, if applicable.",
								Answers: []*model.Answer{
									{Type: "TextArea", QCode: "146", Label: "Comments", Value: "No changes!"},
								},
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

func TestCreateAbs(t *testing.T) {
	test.SetCwdToRoot()

	survey := &model.Survey{
		Title:       "Annual Business Survey",
		SurveyId:    "202",
		FormType:    "1802",
		Respondent:  "12346789012A",
		SubmittedAt: "13 January 2023 16:38:46",
		Sections: []*model.Section{
			{
				Title: "How to complete and reporting period",
				Instances: []*model.Instance{
					{
						Id: 0,
						Questions: []*model.Question{
							{
								Title: "Are you able to report for the period from 01/01/2022 to 31/12/2022?",
								Answers: []*model.Answer{
									{Type: "Radio", QCode: "9999", Label: "label"},
								},
							},
							{
								Title: "What are the dates of the period that you will be reporting for?",
								Answers: []*model.Answer{
									{Type: "Date", QCode: "11", Label: "From", Value: "1/2/2019"},
									{Type: "Date", QCode: "12", Label: "To", Value: "28/3/2019"},
								},
							},
						},
					},
				},
			},
			{
				Title: "Income",
				Instances: []*model.Instance{
					{
						Id: 0,
						Questions: []*model.Question{
							{
								Title: "What is the business's total turnover for the period?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "399", Label: "Total turnover", Value: "56123"},
								},
							},
							{
								Title: "Does your business produce goods or services that protect the environment?",
								Answers: []*model.Answer{
									{Type: "Radio", QCode: "80", Label: "label", Value: "Yes"},
								},
							},
							{
								Title: "Of {total_turnover}, approximately what percentage related to the production of environmental goods or services?",
								Answers: []*model.Answer{
									{Type: "Radio", QCode: "81", Label: "label", Value: "0-9%"},
								},
							},
						},
					},
				},
			},
			{
				Title: "Expenditure",
				Instances: []*model.Instance{
					{
						Id: 0,
						Questions: []*model.Question{
							{
								Title: "What was the business's expenditure on employment costs for the period?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "450", Label: "Total employment costs", Value: "22987"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "403", Label: "Total expenditure on goods and energy products bought for resale", Value: "16723"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure for all other materials, goods and services?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "420", Label: "Total expenditure for all other costs of materials, goods and services", Value: "99883"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure on rates, duties, levies and taxes paid to the government?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "400", Label: "Expenditure on rates, duties, levies and taxes paid to the government", Value: "12345"},
								},
							},
						},
					},
				},
			},
			{
				Title: "Comments",
				Instances: []*model.Instance{
					{
						Id: 0,
						Questions: []*model.Question{
							{
								Title: "Please provide any further details that will help us understand your figures and tell an industry story",
								Answers: []*model.Answer{
									{Type: "TextArea", QCode: "146", Label: "Comments", Value: "My comment!"},
								},
							},
						},
					},
				},
			},
		},
	}

	result := Create(survey)
	err := test.SaveJPG("temp/abs-test.jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for abs_1802 with error: %q", err.Error())
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
						Questions: []*model.Question{
							{
								Title: "What was the business's total expenditure?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "200", Label: "Total expenditure", Value: "10000"},
								},
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
						Questions: []*model.Question{
							{
								Title: "What was the business's expenditure on employment costs for the period?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "450", Label: "Total employment costs", Value: "2000"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "451", Label: "Total expenditure on goods and energy products bought for resale", Value: "1000"},
								},
							},
						},
					},
					{
						Id: 2,
						Questions: []*model.Question{
							{
								Title: "What was the business's expenditure on employment costs for the period?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "450", Label: "Total employment costs", Value: "3000"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "451", Label: "Total expenditure on goods and energy products bought for resale", Value: "1000"},
								},
							},
						},
					},
					{
						Id: 3,
						Questions: []*model.Question{
							{
								Title: "What was the business's expenditure on employment costs for the period?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "450", Label: "Total employment costs", Value: "2000"},
								},
							},
							{
								Title: "During the reporting period, what was your business's expenditure on goods and energy products bought for resale?",
								Answers: []*model.Answer{
									{Type: "Currency", QCode: "451", Label: "Total expenditure on goods and energy products bought for resale", Value: "1000"},
								},
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
