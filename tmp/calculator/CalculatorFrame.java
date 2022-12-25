package tmp.calculator;

import javax.swing.*;
import java.awt.*;
import java.awt.event.*;

public class CalculatorFrame extends JFrame {
    private JTextField textField;
    private JButton[] buttons;
    private JButton buttonAdd;
    private JButton buttonSubtract;
    private JButton buttonMultiply;
    private JButton buttonDivide;
    private JButton buttonEqual;
  
    public CalculatorFrame() {
      setTitle("Calculator");
      setSize(300, 400);
      setDefaultCloseOperation(EXIT_ON_CLOSE);
  
      textField = new JTextField(20);
	  textField.setHorizontalAlignment(JTextField.LEFT);
	  textField.setPreferredSize(new Dimension(200,30));
      add(textField, BorderLayout.NORTH);
	  JPanel panel = new JPanel();
      panel.setLayout(new GridLayout(4,4,1,1));
  
      buttons = new JButton[10];
      for (int i = 0; i < 10; i++) {
		buttons[i] = new JButton("" + i);
		buttons[i].addActionListener(new ActionListener() {
			@Override
			public void actionPerformed(ActionEvent e) {
			  textField.setText(textField.getText() + e.getActionCommand());
			}
		});
		panel.add(buttons[i]);
      }
      buttonAdd = new JButton("+");
      panel.add(buttonAdd);
      buttonSubtract = new JButton("-");
      panel.add(buttonSubtract);
      buttonMultiply = new JButton("*");
      panel.add(buttonMultiply);
      buttonDivide = new JButton("/");
      panel.add(buttonDivide);
      buttonEqual = new JButton("=");
      panel.add(buttonEqual);
	  add(panel, BorderLayout.CENTER);
      setVisible(true);

	  buttonAdd.addActionListener(new ActionListener() {
		@Override
		public void actionPerformed(ActionEvent e) {
		  textField.setText(textField.getText() + "+");
		}
	  });
	  
	  buttonSubtract.addActionListener(new ActionListener() {
		@Override
		public void actionPerformed(ActionEvent e) {
		  textField.setText(textField.getText() + "-");
		}
	  });
	  
	  buttonMultiply.addActionListener(new ActionListener() {
		@Override
		public void actionPerformed(ActionEvent e) {
		  textField.setText(textField.getText() + "*");
		}
	  });
	  
	  buttonDivide.addActionListener(new ActionListener() {
		@Override
		public void actionPerformed(ActionEvent e) {
		  textField.setText(textField.getText() + "/");
		}
	  });
	}
}
  