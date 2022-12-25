package tmp.student;

import java.awt.GridLayout;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextField;

public class Main {
    public static void main(String[] args) {
        JFrame frame = new JFrame("My Window");
        frame.setSize(300, 200);
        frame.setLocation(100, 100);
        
        JPanel panel = new JPanel();
        panel.setLayout(new GridLayout(3, 2));
        
        // 添加组件
        panel.add(new JLabel("Your Name:"));
        JTextField textField = new JTextField(10);
        panel.add(textField);
        JButton button = new JButton("Finished");
        panel.add(button);
        
        frame.add(panel);
        frame.setVisible(true);
        
        // 创建监听器
        ActionListener listener = new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                String name = textField.getText();
                textField.setText("Hello, " + name);
            }
        };
        button.addActionListener(listener);
    }
}
