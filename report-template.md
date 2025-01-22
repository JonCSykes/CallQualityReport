# Customer Service Call Analysis Report
## Overview

**Summary:**
{{ .Overview }}

---

## Strengths

**The representative demonstrated the following strengths:**
{{ range .Strengths }}
- {{ . }}
  {{ end }}

---

## Weaknesses

**Areas for improvement identified during the call:**
{{ range .Weaknesses }}
- {{ . }}
  {{ end }}

---

## Opportunities for Improvement

**Specific suggestions to improve performance in future interactions:**
{{ range .OpportunitiesForImprovement }}
- {{ . }}
  {{ end }}

---

## Performance Ratings

**Performance Criteria Ratings (1 to 5):**

| Criteria                       | Rating |
|-------------------------------|--------|
| Empathy and Emotional Intelligence | {{ .PerformanceRating.EmpathyAndEmotionalIntelligence }} |
| Professionalism               | {{ .PerformanceRating.Professionalism }} |
| Problem-Solving Skills        | {{ .PerformanceRating.ProblemSolvingSkills }} |
| Communication Clarity         | {{ .PerformanceRating.CommunicationClarity }} |
| Customer-Centric Approach     | {{ .PerformanceRating.CustomerCentricApproach }} |
| **Overall Rating**            | **{{ .PerformanceRating.OverallRating }}** |

---

## Conclusion

**Final Thoughts:**
{{ .Conclusion }}

---

**Note:** This report highlights key aspects of the representative's performance and provides actionable feedback to improve customer service quality.
