package filter

func Export(expression Expression) (string, error) {
	var result string
	switch node := expression.(type) {

	case *PrecedenceExpression:
		result += node.String()

	case *NotExpression:
		result += node.String()

	case *LogicalExpression:
		resultLeft, _ := Export(node.Left)
		resultRight, _ := Export(node.Right)
		result += resultLeft + " " + string(node.Operator) + " " + resultRight

	case *ValuePath:
		resultValuePath, _ := Export(node.ValueFilter)
		result += node.AttributePath.String() + "[" + resultValuePath + "]"

	case *AttributeExpression:
		result += node.String()

	default:
	}
	return result, nil
}
